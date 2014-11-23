package dockerclient

import (
	"github.com/fsouza/go-dockerclient"
)

type fakeClient struct{}

func (client *fakeClient) Client() *docker.Client {
	return &docker.Client{}
}
func (client *fakeClient) LatestImageByRegex(regex string) (*docker.APIImages, error) {
	return &docker.APIImages{}, nil
}

func FakeClient() DockerClient {
	return &fakeClient{}
}
