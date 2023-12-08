pipeline {
    agent any
    
    environment {
        registryCredentials = '1bba979e-5f57-4c84-98f9-7e9eca3d7153' // id credentials
        registryUrl = "${REGISTRY_HOST}" // url server registry
        imageName = 'shofwa123/app-golang' // image name
        imageTag = '1.1' // image tag
    }

    stages { // stages
        stage('pull') { // stage
            steps{ // steps
                script { // script
                    docker.image("${imageName}:${imageTag}").pull() // pull image from docker.io
                }
            }
        }
        stage('push') { // stage
            steps { // steps
                script{ // script
                    // Log in to Docker registry
                    docker.withRegistry("https://${registryUrl}", registryCredentials) {
                        // Push the Docker image to the registry
                        docker.image("${imageName}:${imageTag}").push()
                    }
                }
            }
        }
    }
}
