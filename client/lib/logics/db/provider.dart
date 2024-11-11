import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../models/models.dart' as model;

final userProvider = StateProvider<model.User?>((ref) => null);
