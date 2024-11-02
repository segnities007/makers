import 'package:flutter/material.dart';

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
