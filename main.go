package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var instanceID = "i-082e36c19270991b8"

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command (start/stop/exit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "start":
			err := startInstance(instanceID)
			if err != nil {
				fmt.Println("Failed to start instance:", err)
			} else {
				fmt.Println("Instance started.")
			}
		case "stop":
			err := stopInstance(instanceID)
			if err != nil {
				fmt.Println("Failed to stop instance:", err)
			} else {
				fmt.Println("Instance stopped.")
			}
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command. Please use 'start', 'stop', or 'exit'.")
		}
	}
}

func startInstance(instanceID string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewEnvCredentials(),
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
