#!/usr/bin/env ruby
# -*- coding: utf-8 -*-
$LOAD_PATH << __dir__

require 'grpc'
require 'server_pb'
require 'server_services_pb'

class ServerImpl < Server::Transactions::Service
    def add(add_req, _unused_call)
        puts(add_req.inspect)
    end
end

def main
    addr = "0.0.0.0:8080"
    s = GRPC::RpcServer.new
    s.add_http2_port(addr, :this_port_is_insecure)
    s.handle(ServerImpl.new())
    s.run_till_terminated
end

main
