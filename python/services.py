from concurrent import futures
import grpc
import service_pb2
import service_pb2_grpc
from prediction import get_prediction
from temperature import get_temperature


class TemperatureService(service_pb2_grpc.TemperatureServiceServicer):

    def GetTemperature(self, request, context):
        print("Asked temperature")
        return service_pb2.TemperatureResponse(temperature=get_temperature())

    def GetModelResult(self, request, context):
        print("asked model")
        input_string = request.input
        # Process the input string to produce a result
        return service_pb2.ModelResponse(result=get_prediction())

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service_pb2_grpc.add_TemperatureServiceServicer_to_server(TemperatureService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
