package aws

import (
	"deploy/app/conf"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	_ = conf.Init()
	ami1 = "ami-060cdc0f1b2cc0cf2"
	ec2Impl = NewEc2Impl()
)

func TestEc2Impl_shouldReturnNilWhenNoInstanceProvisioned(t *testing.T) {
	assert.Nil(t, ec2Impl.GetByAmi(ami1))
}

func TestEc2Impl_shouldReturnInstanceWhenProvisioned(t *testing.T) {
	defer terminateAll()
	ec2Impl.RunInstance(ami1)
	instances := ec2Impl.GetByAmi(ami1)
	assert.Len(t, instances, 1)

	assert.Equal(t, *instances[0].ImageId, ami1)
	assert.Equal(t, *instances[0].InstanceType, viper.GetString("app.ec2.class"))
}

func TestEc2Impl_WaitInstanceUp(t *testing.T) {
	defer terminateAll()
	instance := ec2Impl.RunInstance(ami1)
	ec2Impl.WaitInstanceUp(instance.Instances[0].InstanceId)
	instances := ec2Impl.GetByAmi(ami1)

	assert.Equal(t, *instances[0].State.Name, "running")
}

func TestEc2Impl_TerminateInstance(t *testing.T) {
	instance := ec2Impl.RunInstance(ami1)
	ec2Impl.TerminateInstance(instance.Instances[0].InstanceId)

	assert.Nil(t, ec2Impl.GetByAmi(ami1))
}

func terminateAll() {
	for _, instance := range ec2Impl.GetByAmi(ami1) {
		ec2Impl.TerminateInstance(instance.InstanceId)
	}
}
