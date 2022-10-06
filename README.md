# gRPC server-side stream 

server side RPC that streams prime factors of a number sent from the
client.

#### To Build:
<code>// clean build </code><br/>
<code>make clean_primes</code><br/>

<code> // build protobuf files </code><br />
<code> make build_protoc</code><br/>

<code>// build client and server </code><br/>
<code>make build_server</code><br/>


<code>make build_client</code><br/>

#### To Run:

<code> ./bin/primes/server </code><br/>

<small>***in a seperate terminal***</small><br/>
<code> ./bin/primes/client </code><br/>