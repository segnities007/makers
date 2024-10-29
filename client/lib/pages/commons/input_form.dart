import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

class MakerInportForm extends HookWidget {
  const MakerInportForm({
    super.key,
    required this.labels,
    required this.validators,
    required this.controllers,
    this.borderRadius = 0,
    this.minWidth = 100,
  });

  final List<String> labels;
  final List<String? Function(String?)> validators;
  final List<TextEditingController> controllers;
  final double borderRadius;
  final double minWidth;

  @override
  Widget build(context) {
    return ConstrainedBox(
      constraints: BoxConstraints(
        minWidth: minWidth,
      ),
      child: Form(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            for (int i = 0; i < labels.length; i++)
              TextFormField(
                validator: validators[i],
                controller: controllers[i],
                decoration: InputDecoration(
                  label: Text(labels[i]),
                  enabledBorder: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(borderRadius),
                  ),
                ),
              ),
          ],
        ),
      ),
    );
  }
}
