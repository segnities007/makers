import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

import './pages/commons/button.dart';

class Splash extends HookWidget {
  const Splash({super.key});

  @override
  Widget build(BuildContext context) {
    useEffect(() {
      WidgetsBinding.instance.addPostFrameCallback((_) {
          context.go("/logins");
      });
      return null;
    }, []);

    return Center(
      child: MakerButton(
        handler: () {
          context.go("/logins");
        },
        label: "go login",
      ),
    );
  }
}
