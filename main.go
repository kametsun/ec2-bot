package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/bwmarrin/discordgo"
)

func main() {
	go startHealthCheckServer()

	Token := os.Getenv("DISCORD_BOT_TOKEN")
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	instanceID := os.Getenv("AWS_INSTANCE_ID")

	if strings.HasPrefix(m.Content, "!start") {
		err := startInstance(instanceID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Failed to start instance: "+err.Error())
		} else {
			s.ChannelMessageSend(m.ChannelID, "Instance started.")
		}
	}

	if strings.HasPrefix(m.Content, "!stop") {
		err := stopInstance(instanceID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Failed to stop instance: "+err.Error())
		} else {
			s.ChannelMessageSend(m.ChannelID, "Instance stopped.")
		}
	}
}

func startInstance(instanceID string) error {
	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
    awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
    awsRegion := os.Getenv("AWS_REGION")

    if awsAccessKeyID == "" || awsSecretAccessKey == "" || awsRegion == "" {
        return fmt.Errorf("AWS credentials or region not set in environment variables")
    }

    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(awsRegion),
        Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
    })
	if err != nil {
		return err
	}

	svc := ec2.New(sess)
	_, err = svc.StartInstances(&ec2.StartInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	})
	if err != nil {
		return err
	}

	fmt.Printf("Instance %s started\n", instanceID)
	return nil
}

func stopInstance(instanceID string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewEnvCredentials(),
	})
	if err != nil {
		return err
	}

	svc := ec2.New(sess)
	_, err = svc.StopInstances(&ec2.StopInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	})
	if err != nil {
		return err
	}

	fmt.Printf("Instance %s stopped\n", instanceID)
	return nil
}

func startHealthCheckServer() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	port := ":8080"
	fmt.Println("Health check server is running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting health check server:", err)
	}
}