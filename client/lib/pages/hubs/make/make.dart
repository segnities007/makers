import 'package:flutter/material.dart';

import '../../commons/input_form.dart';
import '../../commons/button.dart';

class Make extends StatelessWidget {
  Make({super.key});

  final labels = ["Group name", "Description"];
  final validators = [emptyValidator, emptyValidator];
  final controllers = [TextEditingController(), TextEditingController(),];

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          MakerInputForm(padding: 64, labels: labels, validators: validators, controllers: controllers),
          MakerButton(handler: (){}, label: "Make")
        ],
      )
    );
  }
}

String? emptyValidator(String? value){
  return null;
}

