KUBECONFIG="./kubeconfig.yaml"
POD_NAME="goserver-549d7ccccd-jcpql"

apply_resource:
	kubectl apply -f k8s/deployment.yaml -f k8s/service.yaml -f k8s/ingress.yaml --kubeconfig=${KUBECONFIG}

delete_resource:
	kubectl delete -f k8s/deployment.yaml -f k8s/service.yaml -f k8s/ingress.yaml --kubeconfig=${KUBECONFIG}

get_pods:
	kubectl get pods --kubeconfig=${KUBECONFIG}

describe_pod:
	kubectl describe pod ${POD_NAME} --kubeconfig=${KUBECONFIG}

log_pod:
	kubectl logs ${POD_NAME} --kubeconfig=${KUBECONFIG}