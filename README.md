# obsmon
Centralized metadata and monitoring for obscuro testnets

This is an experimental proof-of-concept to see if a service like this would be helpful. This is just the data and monitoring API, the front-end will be in a separate repo.

The goal is to build a dashboard the I find useful for day-to-day testing, monitoring and investigation of the various obscuro testnets.

If other Obscuro contributors find it useful then we might want to productionize it and use it as a central provider of testnet data.

In particular I'd like the service to provide:
- URLs for nodes and tooling
- addresses for relevant network contracts on both the L2 and the L1
- periodic monitoring/aggregation of GH obscuro-test monitoring
