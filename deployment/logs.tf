resource "aws_cloudwatch_log_group" "loxone-weather" {
  name = "loxone-weather"

  tags = {
    Environment = "production"
    Application = "loxone-weather-proxy"
  }
  retention_in_days = 3
}