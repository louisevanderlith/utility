class Service {
  final int hours;
  final int minutes;
  final int seconds;
  final String location;
  final String description;

  Service(this.location, this.description, this.hours, this.minutes, this.seconds);

  Map<String, dynamic> toJson() {
    return {
      "Duration": "${hours}h${minutes}m${seconds}s",
      "Location": location,
      "Description": description,
    };
  }
}
