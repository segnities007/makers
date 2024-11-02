import 'package:flutter/material.dart';

import '../../commons/card.dart';

const double padding = 16;

class Profile extends StatelessWidget {
  const Profile({super.key});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Container(
        padding: const EdgeInsets.all(padding),
        child: const Column(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          MakerCard(isClickable: false ,iconSize: 128, label: "name", description: "description")
        ],
      ),)
    );
  }
}
