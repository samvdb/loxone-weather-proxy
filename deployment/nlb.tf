variable "health_check_interval" {
  default = "30"
}

# The amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused
variable "deregistration_delay" {
  default = "30"
}

resource "aws_lb" "main" {
  name                             = "${var.app}"
  load_balancer_type               = "network"
  enable_cross_zone_load_balancing = "true"

  # launch lbs in public or private subnets based on "internal" variable
  internal = false
  subnet_mapping {
    subnet_id =  "${aws_default_subnet.default_subnet_a.id}"
    allocation_id ="${aws_eip.lb.id}"
  }
}

# adds a tcp listener to the load balancer and allows ingress
resource "aws_lb_listener" "tcp-6066" {
  load_balancer_arn = aws_lb.main.id
  port              = 6066
  protocol          = "TCP"

  default_action {
    target_group_arn = aws_lb_target_group.main-1.arn
    type             = "forward"
  }
}

resource "aws_lb_target_group" "main-1" {
  name                 = "${var.app}"
  port                 = 6066
  protocol             = "TCP"
  vpc_id               = aws_default_vpc.default_vpc.id
  target_type          = "ip"
  deregistration_delay = var.deregistration_delay

  health_check {
    protocol            = "TCP"
    interval            = var.health_check_interval
    healthy_threshold   = 5
    unhealthy_threshold = 5
  }

}