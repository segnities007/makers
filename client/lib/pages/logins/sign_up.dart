import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

import '../commons/input_form.dart';
import '../commons/button.dart';
import '../../logics/http/post.dart';
import '../../logics/db/provider.dart';

class SignUp extends HookConsumerWidget {
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
          label: "create",
          handler: () async {
            try {
              ref.read(userProvider.notifier).state = await postUser(
                email: controllers[0].text,
                password: controllers[1].text,
              );

              final spre = SharedPreferencesAsync();
              await spre.setInt("userid", user.state!.id);

              WidgetsBinding.instance.addPostFrameCallback((_) async {
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
