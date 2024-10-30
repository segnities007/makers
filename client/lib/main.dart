import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

import 'splash.dart';
import './pages/logins/logins.dart';
import './pages/hubs/hubs.dart';

final router = GoRouter(routes: [
  GoRoute(path: "/", builder: (context, state) => const Splash()),
  GoRoute(path: "/logins", builder: (context, state) => const Logins()),
  GoRoute(path: "/hubs", builder: (context, state) => const Hubs())
]);

void main() {
  runApp(const ProviderScope(child: MainApp()));
}

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      routerConfig: router,
    );
  }
}
