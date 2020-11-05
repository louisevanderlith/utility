import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/service.dart';

Future<HttpRequest> createService(Service obj) async {
  var apiroute = getEndpoint("utility");
  var url = "${apiroute}/info";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateService(Key key, Service obj) async {
  var route = getEndpoint("utility");
  var url = "${route}/info/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteService(Key key) async {
  var route = getEndpoint("utility");
  var url = "${route}/info/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
