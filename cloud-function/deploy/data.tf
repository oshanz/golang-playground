data "archive_file" "source" {
  type = "zip"
  source_dir = "${path.module}/../functions"
  output_path = "${path.module}/function.tar.gz"
}