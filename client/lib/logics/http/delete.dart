import 'package:http/http.dart' as http;
import 'package:flutter_dotenv/flutter_dotenv.dart';

Future<bool> deleteUser({
  required int id,
  required int uiid,
}) async {
  final url = Uri.parse("${dotenv.env["API_URL"]}/user").replace(
    queryParameters: {
      "id": id.toString(),
      "uiid": uiid.toString(),
    },
  );

  final res = await http.delete(url);
  if (res.statusCode == 200) {
    return true;
  } else {
    throw Exception("Failed to delete user: ${res.statusCode} ${res.body}");
  }
}
