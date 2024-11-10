import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:image_picker/image_picker.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

import '../../models/models.dart' as model;

Future<model.User> postUser({
  required final String email,
  required final String password,
}) async {
  final url = Uri.parse("${dotenv.env["API_URL"]}/user");
  final res =
      await http.post(url, body: {"email": email, "password": password});
  if (res.statusCode == 200) {
    return model.User.fromJson(jsonDecode(res.body));
  } else {
    throw Exception("Failed to postUser");
  }
}

Future<model.User> postUserInfo({
  required final String name,
  required final int birthday,
  required final String github,
  required final String description,
  required final XFile? image,
  required final List<int> tags,
  required final List<int> groupIDs,
  required final List<int> followIDs,
  required final List<int> directMessageIDs,
}) async {
  final url = Uri.http("localhost:1323", "user");
  final res = await http.post(url, body: {
    "name": name,
    "birthday": birthday,
    "github": github,
    "description": description,
    "tags": tags,
    "groupIDs": groupIDs,
    "followIDs": followIDs,
    "directMessageIDs": directMessageIDs,
  });

  if (res.statusCode == 200) {
    return model.User.fromJson(jsonDecode(res.body));
  } else {
    throw Exception("Failed to postUser");
  }
}
