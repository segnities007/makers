import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import './pages/commons/button.dart';

class Splash extends StatelessWidget {
  const Splash({super.key});

  @override
  Widget build(BuildContext context) {
    return Center(
        child: MakerButton(
      handler: () {
        context.go("/logins");
      },
      label: "go login",
      )
    );
  }
}
