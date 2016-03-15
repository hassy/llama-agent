# Llama Agent

**NOTE:** This software is an extreme alpha. It should work, but you probably shouldn't use it.

## Purpose

This is a companion for [Chaos Llama](https://github.com/hassy/llama-cli).

Llama Agent allows Chaos Llama to interfere with your infrastructure in more ways. By default, Chaos Llama will only terminate EC2 instances, but if you install `llama-agent` on your instances and configure Chaos Llama to talk to it, it can do interesting things like:

- Increasing network latency
- Increase system load
- Kill a random non-root process

## License

MPL 2.0, see [LICENSE.txt](LICENSE.txt) for details.
