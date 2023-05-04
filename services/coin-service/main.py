from concurrent import futures
import time

import grpc

import coin_pb2
import coin_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

class Coiner(coin_pb2_grpc.CoinServicer):

    def GetCoin(self, request, context):
        print("Request: username=%s" % request.username)
        return coin_pb2.CoinResponse(coin='1,000')

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    coin_pb2_grpc.add_CoinServicer_to_server(Coiner(), server)
    server.add_insecure_port('0.0.0.0:50053')
    server.start()

    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()