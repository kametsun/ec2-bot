## Environment Variables

To run this bot, you need to configure the following environment variables:

- **`DISCORD_BOT_TOKEN`**: Your Discord bot token. This token is required to authenticate the bot with the Discord API. You can obtain this token from the [Discord Developer Portal](https://discord.com/developers/applications).

- **`EC2_INSTANCE_ID`**: The ID of the EC2 instance you want to control. This is the unique identifier for your EC2 instance in AWS. You can find this ID in the AWS Management Console under the EC2 section.

- **`AWS_REGION`**: The AWS region where your EC2 instance is located. This should match the region where the EC2 instance was created, e.g., `us-west-2`, `us-east-1`, etc.

### Example

Here’s how to set these environment variables in your terminal:

#### Unix/Linux/MacOS

```bash
export DISCORD_BOT_TOKEN="your-discord-bot-token"
export EC2_INSTANCE_ID="your-ec2-instance-id"
export AWS_REGION="your-aws-region"
