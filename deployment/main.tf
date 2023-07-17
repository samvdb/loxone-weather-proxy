terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "4.45.0"
    }
  }
}
provider "aws" {
  region  = "eu-central-1" #The region where the environment
  #is going to be deployed # Use your own region here
  access_key = var.aws_id # Enter AWS IAM
  secret_key = var.aws_secret_key # Enter AWS IAM
}

resource "aws_ecs_cluster" "my_cluster" {
  name = "loxone-weather-cluster" # Name your cluster here
}


resource "aws_ecs_service" "app_service" {
  name            = "loxone"  # Name the service
  cluster         = "${aws_ecs_cluster.my_cluster.id}"   # Reference the created Cluster
  task_definition = "${aws_ecs_task_definition.loxone-weather-task.arn}" # Reference the task that the service will spin up
  launch_type     = "FARGATE"
  desired_count   = 1 # Set up the number of containers to 3

  load_balancer {
    target_group_arn = "${aws_lb_target_group.main-1.arn}" # Reference the target group
    container_name   = "${aws_ecs_task_definition.loxone-weather-task.family}"
    container_port   = 6066 # Specify the container port
  }

  network_configuration {
    subnets          = ["${aws_default_subnet.default_subnet_a.id}"]
    assign_public_ip = true     # Provide the containers with public IPs
    security_groups  = ["${aws_security_group.nsg_task.id}"] # Set up the security group
  }
}


output "app_url" {
  value = aws_eip.lb.public_ip
}
