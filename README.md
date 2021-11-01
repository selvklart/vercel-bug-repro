# Vercel bug reproduction case

When POSTing a request to a URL backed by a Golang handler func, where the same header is set twice, Vercel seems to crash before calling our handler function.

To reproduce:
1. Deploy this repo to Vercel
2. Call the api/bug.go handler with duplicate HTTP headers: `curl -i -H "Foo: bar" -H "Foo: baz" https://vercel-bug-repro.vercel.app/api/bug.go`
3. Observe the 500 error

Note that if removing the duplicate headers, the request works: `curl -i -H "Food: bar" -H "Foo: baz" https://vercel-bug-repro.vercel.app/api/bug.go`
