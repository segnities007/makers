//基本的には、サーバーから帰ってきたjsonを用いてインスタンスを作成する。

class User {
  final int id;
  final int userInfoID;
  final String email;
  final String password;

  User({
    required this.id,
    required this.email,
    required this.password,
    required this.userInfoID,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json["id"],
      email: json["email"],
      password: json["password"],
      userInfoID: json["userInfoID"],
    );
  }
}

class UserInfo {
  final int id;
  final String name;
  final int birthday;
  final String github;
  final String description;
  final Image? image;
  final List<int> tags;
  final List<int> followIDs;
  final List<int> groupIDs;
  final List<int> dmIDs;

  UserInfo({
    required this.id,
    this.name = "",
    this.birthday = -1,
    this.github = "",
    this.description = "",
    this.image,
    this.tags = const [],
    this.followIDs = const [],
    this.groupIDs = const [],
    this.dmIDs = const [],
  });

  factory UserInfo.fromJson(Map<String, dynamic> json) {
    return UserInfo(
      id: json["id"],
      name: json["name"],
      birthday: json["birthday"],
      github: json["github"],
      description: json["description"],
      image: json["image"],
      tags: json["tags"],
      followIDs: json["followIDs"],
      groupIDs: json["groupIDs"],
      dmIDs: json["dmIDs"],
    );
  }
}

class Image {
  final int id;
  final int userID;
  final String url;

  Image({
    required this.id,
    required this.userID,
    required this.url,
  });

  factory Image.fromJson(Map<String, dynamic> json) {
    return Image(
      id: json["id"],
      userID: json["userID"],
      url: json["url"],
    );
  }
}

class Group {
  final int id;
  final String name;
  final List<int> userIDs;
  final List<int> messageIDs;

  Group({
      required this.id,
      required this.name,
      required this.userIDs,
      required this.messageIDs
      });

  factory Group.fromJson(Map<String, dynamic> json) {
    return Group(
      id: json["id"],
      name: json["json"],
      userIDs: json["userIDs"],
      messageIDs: json["messageIDs"],
    );
  }
}

class Message {
  final int id;
  final int userID;
  final String message;

  Message({
    required this.id,
    required this.userID,
    required this.message,
  });

  factory Message.fromJson(Map<String, dynamic> json) {
    return Message(
      id: json["id"],
      userID: json["userID"],
      message: json["message"],
    );
  }
}

class DirectMessage {
  final int id;
  final int firstID;
  final int secondID;
  final List<int> messageIDs;

  DirectMessage({
    required this.id,
    required this.firstID,
    required this.secondID,
    required this.messageIDs,
  });

  factory DirectMessage.fromJson(Map<String, dynamic> json) {
    return DirectMessage(
      id: json["id"],
      firstID: json["firstID"],
      secondID: json["secondID"],
      messageIDs: json["messageIDs"],
    );
  }
}

class Tags {
  final int id;
  final int tagID;
  final List<int> userIDs;
  final List<int> groupIDs;

  Tags({
    required this.id,
    required this.tagID,
    required this.userIDs,
    required this.groupIDs,
  });
}
