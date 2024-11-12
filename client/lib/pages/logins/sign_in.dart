import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../../logics/db/share_preference.dart';
import '../../logics/http/get.dart';
import '../commons/input_form.dart';
import '../commons/button.dart';
import '../../logics/db/provider.dart';

class SignIn extends HookConsumerWidget {
  SignIn({super.key});

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
  Widget build(context, ref) {
    final user = ref.watch(userProvider.notifier);

    return Center(
        child: Column(children: [
      MakerInputForm(
          padding: 64,
          labels: labels,
          validators: validators,
          controllers: controllers),
      MakerButton(
          label: "In",
          handler: () async {
            try {
              ref.read(userProvider.notifier).state = await getUserWithEP(
                  email: controllers[0].text, password: controllers[1].text);
              await SharedPreferencesAsync().setInt("userid", user.state!.id);
              WidgetsBinding.instance.addPostFrameCallback((_) {
                context.go("/hubs");
              });
            } catch (e) {
              print(e);
            }
          })
    ]));
  }
}

String? emptyValidator(value) {
  return null;
}
