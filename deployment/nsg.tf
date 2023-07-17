resource "aws_security_group" "nsg_task" {
  name        = "${var.app}-task"
  description = "Limit connections from internal resources while allowing ${var.app}-task to connect to all external resources"
  vpc_id      = aws_default_vpc.default_vpc.id
}

# Rules for the TASK (Targets the LB's IPs)
resource "aws_security_group_rule" "nsg_task_ingress_rule" {
  description = "Only allow connections from the NLB on port 6066"
  type        = "ingress"
  from_port   = 6066
  to_port     = 6066
  protocol    = "tcp"
  cidr_blocks = formatlist(
    "%s/32",
    flatten(data.aws_network_interface.nlb.*.private_ips),
  )

  security_group_id = aws_security_group.nsg_task.id
}

resource "aws_security_group_rule" "nsg_task_egress_rule" {
  description = "Allows task to establish connections to all resources"
  type        = "egress"
  from_port   = "0"
  to_port     = "0"
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.nsg_task.id
}

# lookup the ENIs associated with the NLB
data "aws_network_interface" "nlb" {


  filter {
    name   = "description"
    values = ["ELB ${aws_lb.main.arn_suffix}"]
  }

  filter {
    name = "subnet-id"
    values = [aws_default_subnet.default_subnet_a.id]
  }
}