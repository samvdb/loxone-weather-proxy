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
  name            = "app-first-service"  # Name the service
  cluster         = "${aws_ecs_cluster.my_cluster.id}"   # Reference the created Cluster
  task_definition = "${aws_ecs_task_definition.loxone-weather-task.arn}" # Reference the task that the service will spin up
  launch_type     = "FARGATE"
  desired_count   = 1 # Set up the number of containers to 3

  load_balancer {
    target_group_arn = "${aws_lb_target_group.target_group-1.arn}" # Reference the target group
    container_name   = "${aws_ecs_task_definition.loxone-weather-task.family}"
    container_port   = 6066 # Specify the container port
  }

  network_configuration {
    subnets          = ["${aws_default_subnet.default_subnet_a.id}", "${aws_default_subnet.default_subnet_b.id}"]
    assign_public_ip = true     # Provide the containers with public IPs
    security_groups  = ["${aws_security_group.service_security_group.id}"] # Set up the security group
  }
}

resource "aws_security_group" "service_security_group" {
  ingress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    # Only allowing traffic in from the load balancer security group
    security_groups = ["${aws_security_group.load_balancer_security_group.id}"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

output "app_url" {
  value = aws_alb.application_load_balancer.dns_name
}
#output "app_ip" {
#  value = aws_alb.application_load_balancer.
#}