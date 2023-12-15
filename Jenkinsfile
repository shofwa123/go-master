pipeline {
    agent any
    
    environment {
        registryCredentials = 'b8ebe5ff-192f-49fe-90e0-47b187cb1bf2' // id credentials
        registryUrl = "${REGISTRY_HOST}" // url server registry
        k8sFile = "k8s-dev.yaml"
        k8sCredentials = credentials("1bba979e-5f57-4c84-98f9-7e9eca3d7153")
        k8sNS="k8s-crud-golang"
        k8sCA="${K8SDEV_CA_64e}".bytes('UTF-8').decodeBase64().toString('UTF-8')
        imageName = 'shofwa123/app-golang' // image name
        imageTag = '1.1' // image tag
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
                    withKubeConfig(caCertificate: '''-----BEGIN CERTIFICATE-----
                    MIIDBTCCAe2gAwIBAgIIdAkEUuUKGg0wDQYJKoZIhvcNAQELBQAwFTETMBEGA1UE
                    AxMKa3ViZXJuZXRlczAeFw0yMzExMjIwNzA4NDRaFw0zMzExMTkwNzEzNDRaMBUx
                    EzARBgNVBAMTCmt1YmVybmV0ZXMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
                    AoIBAQDfn5UZ4Kq0SrqxZ8Dk3aGbyMl8MgVfQyjfgHkHQ4pfDjWg9whkwJwHNu4r
                    UZRDKTWR0iVuYSdVUtHfmYIIkCHkqrqS2Xo7jhgaKoYbB2cTU/QH0oBNytL940Na
                    411wZR5Mf14B7Gmi7tNEm3xuc/tHR5HLKF0yLsSjo6SzGC6xLiRjB2fldiSY98Nm
                    LJ1OU5HJuWfcvMKn49xbo65+xzyzRvIsVc6mme4GCZhaN+fEtmn9UUCo7+0Zw0CB
                    XIo7r9pYgiZzFECqPitS6+Yo16WKi/5sh7AgadfJNZNPCcKFeMUqwuFGIFlh0p11
                    SFsMBdZqQ1Pprst/r/W22++bHdKJAgMBAAGjWTBXMA4GA1UdDwEB/wQEAwICpDAP
                    BgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBSq/W+NRLfMkRUXxPOu4PIKaIXuCzAV
                    BgNVHREEDjAMggprdWJlcm5ldGVzMA0GCSqGSIb3DQEBCwUAA4IBAQDP2AhXoWNE
                    xM2ABUbdr1JHvyMWL4fqZLCrybjk5OHX0n6/Cf6yqoYAnXInGj7lTtYI50cA4ny8
                    esYRQkCuefgb8KOlrjJLw7Tin1qyxZZ61mTl7Y5IKgDSQkjDpoXhQQ5zUPePSa/x
                    nAGuygWpeRx9TpBL8vYBLBUnCjmmSzIaPxevt0bU0VnpBLQvw6Xz7qz9ITUdi83C
                    cH+T9Boaq/lZ41BYJIUzXjT++uhsnqcRgBhPqc2ZOKv8HCQa1kYq/BIAeZBsRYEE
                    kkmQJp4IvJpr3Q18cqCMlq8NQ3CTzeXrM6VqRmQVAG76CE22SFrXXju8Jf7hP59O
                    NgIyQpYUaoOf
                    -----END CERTIFICATE-----''', clusterName: '', contextName: '', credentialsId: 'a5de36be-7020-45cf-83dc-fbffb95d0c60', namespace: 'k8s-crud-golang', restrictKubeConfigAccess: false, serverUrl: 'https://10.26.11.27:6443') {
                        sh "kubectl get pod -n k8s-crud-golang"
                    }
                }
            }
        }
    }

}
