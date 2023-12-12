pipeline {
    agent any
    
    environment {
        registryCredentials = '9afe98fb-3fc9-4611-99f9-377371d39c1f' // id credentials
        registryUrl = "${REGISTRY_HOST}" // url server registry
        k8sFile = "k8s-dev.yaml"
        k8sCredentials = credentials("1e74ca30-488f-4406-b794-e6d913da103e")
        k8sNS="k8s-rbac"
        imageName = 'nginx' // image name
        imageTag = 'latest' // image tag
    }

    stages { // stages
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('build') { // stage
            steps{ // steps
                script { // script
                    sh "docker build -t ${imageName}:${imageTag} ."
                }
            }
        }
        stage('push') { // stage ci continues integration
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
        stage('continous deployment') {
            steps {
                script {
                    // Set KUBECONFIG environment variable
                    withKubeConfig([
                        serverUrl:"https://10.26.11.27:6443", 
                        caCertificate:"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURCVENDQWUyZ0F3SUJBZ0lJZEFrRVV1VUtHZzB3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TXpFeE1qSXdOekE0TkRSYUZ3MHpNekV4TVRrd056RXpORFJhTUJVeApFekFSQmdOVkJBTVRDbXQxWW1WeWJtVjBaWE13Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLCkFvSUJBUURmbjVVWjRLcTBTcnF4WjhEazNhR2J5TWw4TWdWZlF5amZnSGtIUTRwZkRqV2c5d2hrd0p3SE51NHIKVVpSREtUV1IwaVZ1WVNkVlV0SGZtWUlJa0NIa3FycVMyWG83amhnYUtvWWJCMmNUVS9RSDBvQk55dEw5NDBOYQo0MTF3WlI1TWYxNEI3R21pN3RORW0zeHVjL3RIUjVITEtGMHlMc1NqbzZTekdDNnhMaVJqQjJmbGRpU1k5OE5tCkxKMU9VNUhKdVdmY3ZNS240OXhibzY1K3h6eXpSdklzVmM2bW1lNEdDWmhhTitmRXRtbjlVVUNvNyswWncwQ0IKWElvN3I5cFlnaVp6RkVDcVBpdFM2K1lvMTZXS2kvNXNoN0FnYWRmSk5aTlBDY0tGZU1VcXd1RkdJRmxoMHAxMQpTRnNNQmRacVExUHByc3Qvci9XMjIrK2JIZEtKQWdNQkFBR2pXVEJYTUE0R0ExVWREd0VCL3dRRUF3SUNwREFQCkJnTlZIUk1CQWY4RUJUQURBUUgvTUIwR0ExVWREZ1FXQkJTcS9XK05STGZNa1JVWHhQT3U0UElLYUlYdUN6QVYKQmdOVkhSRUVEakFNZ2dwcmRXSmxjbTVsZEdWek1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRRFAyQWhYb1dORQp4TTJBQlViZHIxSkh2eU1XTDRmcVpMQ3J5YmprNU9IWDBuNi9DZjZ5cW9ZQW5YSW5HajdsVHRZSTUwY0E0bnk4CmVzWVJRa0N1ZWZnYjhLT2xyakpMdzdUaW4xcXl4Wlo2MW1UbDdZNUlLZ0RTUWtqRHBvWGhRUTV6VVBlUFNhL3gKbkFHdXlnV3BlUng5VHBCTDh2WUJMQlVuQ2ptbVN6SWFQeGV2dDBiVTBWbnBCTFF2dzZYejdxejlJVFVkaTgzQwpjSCtUOUJvYXEvbFo0MUJZSklVelhqVCsrdWhzbnFjUmdCaFBxYzJaT0t2OEhDUWExa1lxL0JJQWVaQnNSWUVFCmtrbVFKcDRJdkpwcjNRMThjcUNNbHE4TlEzQ1R6ZVhyTTZWcVJtUVZBRzc2Q0UyMlNGclhYanU4SmY3aFA1OU8KTmdJeVFwWVVhb09mCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K", 
                        credentialsId: ${k8sCredentials}, 
                        namespace: ${k8sNS}
                        ]) {
                        // command
                        sh "kubectl get pod -n k8s-rbac"
                    }
                }
            }
        }
    }

}