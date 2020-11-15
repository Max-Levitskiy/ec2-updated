package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/viper"
)

type Ec2 interface {
	GetByAmi(ami string) []*ec2.Instance
	RunInstance(ami string) *ec2.Reservation
	WaitInstanceUp(instanceId *string)
	TerminateInstance(instanceId *string)
}

type Ec2Impl struct {
	sess *session.Session
	ec2 *ec2.EC2
}

func NewEc2Impl() *Ec2Impl {
	e := &Ec2Impl{}
	if s, err := session.NewSession(&aws.Config{Region: aws.String(viper.GetString("app.region"))}); err == nil {
		e.sess = s
	} else {
		panic(err)
	}
	e.ec2 = ec2.New(e.sess)
	return e
}

func (e *Ec2Impl) GetByAmi(ami string) []*ec2.Instance {
	var insts []*ec2.Instance
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			// TODO: find AMI parameter name
			//{
			//	Name:   aws.String("product-code"),
			//	Values: []*string{&ami},
			//},
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running")},
			},

		},
	}
	if instances, err := e.ec2.DescribeInstances(input); err == nil {
		for _, reservation := range instances.Reservations {
			for _, instance := range reservation.Instances {
				if *instance.ImageId == ami {
					insts = append(insts, instance)
				}
			}
		}
	} else {
		panic(err)
	}
	return insts
}

func (e *Ec2Impl) RunInstance(ami string) *ec2.Reservation {
	ec2Class := viper.GetString("app.ec2.class")
	input := &ec2.RunInstancesInput{
		ImageId:      &ami,
		InstanceType: &ec2Class,
		MaxCount:     aws.Int64(1),
		MinCount:     aws.Int64(1),
	}
	if instances, err := e.ec2.RunInstances(input); err == nil {
		return instances
	} else {
		panic(err)
	}
	return nil
}

func (e *Ec2Impl) WaitInstanceUp(instanceId *string) {
	input := &ec2.DescribeInstanceStatusInput{
		InstanceIds: []*string{instanceId},
	}
	if err := e.ec2.WaitUntilInstanceStatusOk(input); err != nil {
		panic(err)
	}
}

func (e *Ec2Impl) TerminateInstance(instanceId *string) {
	_, _ = e.ec2.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: []*string{instanceId},
	})
}
