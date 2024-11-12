import 'package:http/http.dart' as http;
import 'package:flutter_dotenv/flutter_dotenv.dart';
import '../../models/models.dart' as model;
import 'dart:convert';

Future<model.User> getUser({
  required final int id,
}) async {

  final url = Uri.parse("${dotenv.env["API_URL"]}/user/$id");
  final res = await http.get(url);

  if (res.statusCode == 200) {

    return model.User.fromJson(jsonDecode(res.body));
  } else {
    throw Exception("Failed to getUser");
  }
}

Future<model.User> getUserWithEP({
  required final String email,
  required final String password,
}) async {
  final url = Uri.parse("${dotenv.env["API_URL"]}/user").replace(
    queryParameters: {
      "email": email,
      "password": password,
    },
  );

  final res = await http.get(url);

  if (res.statusCode == 200) {
    return model.User.fromJson(jsonDecode(res.body));
  } else {
    throw Exception("Failed to getUser");
  }
}
