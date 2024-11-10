import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../commons/input_form.dart';
import '../commons/button.dart';
import '../../logics/http/post.dart';
import '../../models/models.dart' as model;

class SignUp extends HookWidget {
  SignUp({super.key});

  final List<String> labels = ["Email", "Password"];
  final List<String? Function(String?)> validators = [
    emptyValidator,
    emptyValidator,
  ];
  final List<TextEditingController> controllers = [
    TextEditingController(),
    TextEditingController(),
  ];

  @override
  Widget build(BuildContext context) {
    final user = useState<model.User?>(null);

    return Center(
        child: Column(children: [
      (user.value != null)
          ? Text("${user.value!.id} ${user.value!.email}")
          : const Text("null"),
      MakerInputForm(
          padding: 64,
          labels: labels,
          validators: validators,
          controllers: controllers),
      MakerButton(
          label: "create",
          handler: () async {
            try {
              user.value = await postUser(
                email: controllers[0].text,
                password: controllers[1].text,
              );

              final spre = SharedPreferencesAsync();
              await spre.setInt("userid", user.value!.id);

              WidgetsBinding.instance.addPostFrameCallback((_) {
                context.go("/hubs");
              });
            } catch (e) {
              //TODO change print
              print(e);
            }
          })
    ]));
  }
}

String? emptyValidator(value) {
  return null;
}
