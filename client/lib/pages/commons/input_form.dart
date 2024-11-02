import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';



class MakerInputForm extends HookWidget {
  const MakerInputForm({
    super.key,
    required this.labels,
    required this.validators,
    required this.controllers,
    this.borderRadius = 8,
    this.maxWidth = 1000,
    this.padding = 32,
  });

  final List<String> labels;
  final List<String? Function(String?)> validators;
  final List<TextEditingController> controllers;
  final double borderRadius;
  final double maxWidth;
  final double padding;


  @override
  Widget build(context) {
    return Form(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          for (int i = 0; i < labels.length; i++)
            Padding(
              padding: EdgeInsets.fromLTRB(padding, padding/2, padding, padding/2),
              child: ConstrainedBox(
                constraints: BoxConstraints(
                  maxWidth: maxWidth,
                ),
                child: TextFormField(
                  validator: validators[i],
                  controller: controllers[i],
                  decoration: InputDecoration(
                    filled: true,
                    fillColor: Colors.white,
                    label: Text(labels[i]),
                    enabledBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(borderRadius),
                    ),
                  ),
                ),
              ),
            )
        ],
      ),
    );
  }
}
