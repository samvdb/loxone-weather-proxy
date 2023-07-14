resource "aws_ecs_task_definition" "loxone-weather-task" {
  family                   = "loxone-weather-proxy" # Name your task
  container_definitions    = <<DEFINITION
  [
    {
      "name": "loxone-weather-proxy",
      "image": "ghcr.io/samvdb/loxone-weather-proxy:v1.0.2",
      "essential": true,
      "portMappings": [
        {
          "containerPort": 6066,
          "hostPort": 6066
        }
      ],
      "environment": [
                {
                    "name": "TOMORROW_APIKEY",
                    "value": "${var.tomorrow_api_key}"
                }
            ],
      "memory": 512,
      "cpu": 256
    }
  ]
  DEFINITION
  requires_compatibilities = ["FARGATE"] # use Fargate as the launch type
  network_mode             = "awsvpc"    # add the AWS VPN network mode as this is required for Fargate
  memory                   = 512         # Specify the memory the container requires
  cpu                      = 256         # Specify the CPU the container requires
  execution_role_arn       = "${aws_iam_role.ecsTaskExecutionRole.arn}"
}