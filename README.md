# go-bridge
Go Bridge Pattern


The code defines three interfaces:

CollectorInterface - Represents the collector component
and has the Collect method.

ResolverInterface - Represents the resolver component,
which depends on the collector, and has the Resolve
and SetCollector methods.

DispatcherInterface - Represents the dispatcher component,
which depends on the resolver, and has the Dispatch
and SetResolver methods.

The Bridge pattern is applied to the connection between the dispatcher and the
resolver and between the resolver and the collector.
The code provides concrete implementations for these interfaces:

ConcreteCollector: Implements the CollectorInterface.

ConcreteResolver: Implements the ResolverInterface and depends
on the CollectorInterface via the SetCollector method.

ConcreteDispatcher: Implements the DispatcherInterface and depends
on the ResolverInterface via the SetResolver method.

In the main function, instances of the concrete implementations are created,
and the dependencies are set up using the SetCollector and SetResolver methods.
Finally, the Dispatch method of the ConcreteDispatcher is called.