import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import '../commons/input_form.dart';
import '../commons/button.dart';

class SignUp extends StatelessWidget {
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
    return Center(
        child: Column(children: [
        MakerInportForm(
            labels: labels, validators: validators, controllers: controllers),
      MakerButton(
          label: "create",
          handler: () {
            // TODO
            WidgetsBinding.instance.addPostFrameCallback((_) {
              context.go("/hubs");
            });
          })
    ]));
  }
}

String? emptyValidator(value) {
  return null;
}
