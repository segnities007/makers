import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

import '../../commons/input_form.dart';
import '../../commons/button.dart';
import '../../commons/card.dart';

const double padding = 32;

class Search extends StatelessWidget {
  Search({super.key});

  final List<String> labels = ["Search"];
  final List<String? Function(String?)> validators = [
    emptyValidator,
  ];
  final List<TextEditingController> controllers = [
    TextEditingController(),
  ];

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: const EdgeInsets.fromLTRB(padding, 8, padding, padding),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Row(
              children: [
                Expanded(
                  child: MakerInportForm(
                    labels: labels,
                    validators: validators,
                    controllers: controllers,
                    maxWidth: double.infinity,
                  ),
                ),
                const SizedBox(width: 8),
                MakerButton(
                  label: "Enter",
                  height: 0.5,
                  handler: () {
                    //TODO
                  },
                ),
              ]
            ),
            const Expanded(
              child: GroupList(),
            ),
          ],
        ));
  }
}

class GroupList extends HookWidget {
  const GroupList({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.green[50],
        borderRadius: BorderRadius.circular(16),
      ),
      child: ListView.builder(
        itemCount: 100,
        itemBuilder: (context, index) => MakerCard(index: index),
      ),
    );
  }
}

String? emptyValidator(String? value) {
  return null;
}
