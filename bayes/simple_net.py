import bayespy as bp
mu = pb.nodes.Normal(0,1e-3)
tau = bp.nodes.Gamma(1e-3, 1e-3)
y = bp.nodes.Normal(mu, tau, plates=(10,))

y.observe(data)