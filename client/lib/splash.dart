import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'logics/db/provider.dart';
import 'logics/http/get.dart';

class Splash extends HookConsumerWidget {
  const Splash({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final user = ref.watch(userProvider.notifier);

    useEffect(() {
      WidgetsBinding.instance.addPostFrameCallback((_) async {
        try {
          final id = await SharedPreferencesAsync().getInt("userid");
          if (id != null) {
            user.state = await getUser(id: id);
            context.go("/hubs");
          } else {
            context.go("/logins");
          }
        } catch (e) {
          print("Error: $e");
          context.go("/logins");
        }
      });
      return null;
    }, []);

    return const Center(child: Text("hello world"));
  }
}
