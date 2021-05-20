# Server plugin: NodeAttestor "join_token"

*Must be used in conjunction with the agent-side join_token plugin*

The `join_token` plugin attests a node based on a pre-shared, one-time-use join token. A
token must be generated by the server before it can be used to attest a node.

The server uses the token to generate a SPIFFE ID with the form:

```
spiffe://<trust domain>/spire/agent/join_token/<token>
```

This plugin has no configuration options. Tokens may be generated through the CLI utility
(`spire-server token generate`) or through the registration API.