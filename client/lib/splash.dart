import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

import 'logics/db/provider.dart';
import 'logics/http/get.dart';
// import 'models/models.dart' as model;

class Splash extends HookConsumerWidget {
  const Splash({super.key});

  @override
  Widget build(context, ref) {
    final user = ref.watch(userProvider.notifier);

    useEffect(() {
      int? id;
      WidgetsBinding.instance.addPostFrameCallback((_) async {
        final spre = SharedPreferencesAsync();
        id = await spre.getInt("userid");
        if (id != null) {
          user.state = await getUser(id: id as int);
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
