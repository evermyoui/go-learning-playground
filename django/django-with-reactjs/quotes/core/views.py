from rest_framework.views import APIView
from rest_framework.response import Response
from .models import React
from .serializer import ReactSerializer

# Create your views here.

class ReactView(APIView):
    serializer_class = ReactSerializer

    def get (self, request):
        detail = [
            {
                "name": obj.name,
                "detail": obj.detail 
            }
            for obj in React.objects.all()
        ]
        return Response(detail)
    
    def post (self, request):
        serializer = ReactSerializer(data=request.data)
        if serializer.is_valid(raise_exception=True):
            serializer.save()
            return Response(serializer.data)