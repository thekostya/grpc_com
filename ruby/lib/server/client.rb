#!/usr/bin/env ruby
# -*- coding: utf-8 -*-
$LOAD_PATH << __dir__

require 'grpc'
require 'server_pb'
require 'server_services_pb'

def main
    stub = Server::Transactions::Stub.new('localhost:8888', :this_channel_is_insecure)
    transaction = Server::NewTransaction.new(
        env: Server::NewTransaction::Env::IL,
        company: Server::Company.new(id: 123, name: "test company name"),
        transaction: Server::Transaction.new(id: "432", amount: 78324),
        user: Server::User.new(username: "test username"),
        order: Server::Order.new(id: 3455),
        calculation: Server::Calculation.new(id: 562),
    )
    resp = stub.add(transaction)
end

main
