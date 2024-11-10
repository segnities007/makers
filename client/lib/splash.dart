import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:shared_preferences/shared_preferences.dart';

class Splash extends HookWidget {
  const Splash({super.key});

  @override
  Widget build(context) {
    useEffect(() {
      int? id;
      WidgetsBinding.instance.addPostFrameCallback((_) async {
        final spre = SharedPreferencesAsync();
        id = await spre.getInt("userid");
        if (id != null) {
          context.go("/hubs");
        } else {
          context.go("/logins");
        }
      });
      return null;
    }, []);

    return const Center(child: Text("hello world"));
  }
}
