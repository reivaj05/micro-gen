# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.http import HttpResponse, JsonResponse
from django.views.decorators.csrf import csrf_exempt

import json


# TODO: Do not leave csrf decorator here
@csrf_exempt
def app_list(request):
    if request.method == 'GET':
        print('TODO: Implement list')
        return JsonResponse({'data': []}, status=200)

    elif request.method == 'POST':
        print('TODO: Implement post')
        try:
            data = json.loads(request.body)
            if is_data_valid(data):
                return JsonResponse({}, status=201)
        except ValueError:
            pass
        return JsonResponse({'error': 'Bad request'}, status=400)
    return HttpResponse(status=404)


@csrf_exempt
def app_detail(request, pk):
    if request.method == 'GET':
        print('TODO: Implement get', pk)
        return JsonResponse({}, status=200)

    elif request.method == 'PUT':
        print('TODO: Implement put', pk)
        try:
            data = json.loads(request.body)
            if is_data_valid(data):
                return JsonResponse({}, status=200)
        except ValueError:
            pass
        return JsonResponse({'error': 'Bad request'}, status=400)

    elif request.method == 'DELETE':
        print('TODO: Implement delete', pk)
        return HttpResponse(status=200)
    return HttpResponse(status=404)


def is_data_valid(data):
    # TODO: Implement your own logic
    return True
