class Todo {
  String? sId;
  String? title;
  String? description;
  bool? completed;
  String? createdAt;
  String? updatedAt;

  Todo(
      {this.sId,
      this.title,
      this.description,
      this.completed,
      this.createdAt,
      this.updatedAt});

  Todo.fromJson(Map<String, dynamic> json) {
    sId = json['_id'];
    title = json['title'];
    description = json['description'];
    completed = json['completed'];
    createdAt = json['createdAt'];
    updatedAt = json['updatedAt'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['_id'] = sId;
    data['title'] = title;
    data['description'] = description;
    data['completed'] = completed;
    data['createdAt'] = createdAt;
    data['updatedAt'] = updatedAt;
    return data;
  }
}
