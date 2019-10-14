## go-nsq

[![Build Status](https://secure.travis-ci.org/nsqio/go-nsq.svg?branch=master)][travis] [![GoDoc](https://godoc.org/github.com/nsqio/go-nsq?status.svg)](https://godoc.org/github.com/nsqio/go-nsq) [![GitHub release](https://img.shields.io/github/release/nsqio/go-nsq.svg)](https://github.com/nsqio/go-nsq/releases/latest)

The official Go package for [NSQ][nsq].

### Docs

See [godoc][nsq_gopkgdoc] and the [main repo apps][apps] directory for examples of clients built
using this package.

### Tests

Tests are run via `./test.sh` (which requires `nsqd` and `nsqlookupd` to be installed).

[nsq]: https://github.com/nsqio/nsq
[nsq_gopkgdoc]: http://godoc.org/github.com/nsqio/go-nsq
[apps]: https://github.com/nsqio/nsq/tree/master/apps
[travis]: http://travis-ci.org/nsqio/go-nsq


消息中间件本质上就是一种很简单的数据结构——队列

需要考虑性能、容灾、可靠性等等因素

为什么需要一些中间件

比如网购了一件东西后
1. 消息通知系统，通知商家，有新订单，及时处理。
    
