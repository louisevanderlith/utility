class Service {
  final String url;

  Service(this.url);

  Map<String, dynamic> toJson() {
    return {"Url": url};
  }
}
