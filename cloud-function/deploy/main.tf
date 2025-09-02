terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.34.0"
    }
  }
}

provider "google" {
  project     = var.project
}

resource "random_id" "default" {
  byte_length = 4
}

resource "google_storage_bucket" "function_bucket" {
  name = "${random_id.default.hex}-fucntion-source"
  location = var.location
}

resource "google_storage_bucket_object" "bucket-object" {
  name = "hello-function-source.zip"
  source = data.archive_file.source.output_path
  bucket = google_storage_bucket.function_bucket.name
}

resource "google_cloudfunctions2_function" "hello-function" {
  name = "hello"
  location = var.location

  service_config {
    max_instance_count = 1
    timeout_seconds = 5
    available_cpu = 0.08
    available_memory = "128Mi"
  }

  build_config {
    runtime = "go125"
    entry_point = "hello"
    source {
      storage_source {
        bucket = google_storage_bucket.function_bucket.name
        object = google_storage_bucket_object.bucket-object.name
      }
    }
  }
}

resource "google_cloud_run_service_iam_member" "member" {
  service = google_cloudfunctions2_function.hello-function.name
  location = google_cloudfunctions2_function.hello-function.location
  member = "allUsers"
  role = "roles/run.invoker"
}