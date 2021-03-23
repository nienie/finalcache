package registry

// CopyService make a copy of service
func CopyService(service *Service) *Service {
	// copy service
	s := new(Service)
	*s = *service

	// copy nodes
	nodes := make([]*Node, len(service.Nodes))
	for j, node := range service.Nodes {
		n := new(Node)
		*n = *node
		nodes[j] = n
	}
	s.Nodes = nodes

	// copy endpoints
	eps := make([]*Endpoint, len(service.Endpoints))
	for j, ep := range service.Endpoints {
		e := new(Endpoint)
		*e = *ep
		eps[j] = e
	}
	s.Endpoints = eps
	return s
}

// Copy makes a copy of services
func Copy(current []*Service) []*Service {
	services := make([]*Service, len(current))
	for i, service := range current {
		services[i] = CopyService(service)
	}
	return services
}
