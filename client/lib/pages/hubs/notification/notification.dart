import 'package:flutter/material.dart';
// import '../../../models/test.dart';

import '../../commons/card.dart';

const double radius = 16;
const double padding = 16;

class Notification extends StatelessWidget {
  const Notification({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(padding),
      child: Container(
        padding: const EdgeInsets.all(padding),
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(radius),
          color: Colors.green[100],
        ),
        child: const Column(
          // mainAxisAlignment: MainAxisAlignment.center,
          children: [
            MakerCard(label: "notify", description: "This is a notify."),
          ],
        ),
      ),
    );
  }
}

class NotifyTile extends StatelessWidget {
  const NotifyTile({
    super.key,
    required this.icon,
    required this.title,
    required this.description,
  });

  final IconData icon;
  final String title;
  final String description;

  @override
  Widget build(context) {
    return Material(
        child: InkWell(
            onTap: () {},
            child: const Row(
                mainAxisAlignment: MainAxisAlignment.start, children: [])));
  }
}
