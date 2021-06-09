module Fastlane
  module Actions
    module SharedValues
      BUILD_GOLANG_ANDROID_CUSTOM_VALUE = :BUILD_GOLANG_ANDROID_CUSTOM_VALUE
    end

    class BuildGolangAndroidAction < Action
      def self.run(params)
        # fastlane will take care of reading in the parameter and fetching the environment variable:
        #UI.message "Parameter API Token: #{params[:api_token]}"

        if !File.directory?("build/")
          sh "mkdir build/"
        end
        sh "cd build/"
        sh "go get -u gioui.org/cmd/gogio"
        sh "gogio -target android -o build/PortScanner-android.apk -icon assets/icon.png -version=4 -appid com.portscanner.superredstone ."

        # Actions.lane_context[SharedValues::BUILD_GOLANG_ANDROID_CUSTOM_VALUE] = "my_val"
      end

      #####################################################
      # @!group Documentation
      #####################################################

      def self.description
        "Builds Gioui for android"
      end

      def self.details
        "Builds Gioui for android"
      end

      def self.available_options
        # Define all options your action supports.

        # Below a few examples
        [

          FastlaneCore::ConfigItem.new(key: :api_token,
                                       env_name: "FL_BUILD_GOLANG_ANDROID_API_TOKEN", # The name of the environment variable
                                       description: "API Token for BuildGolangAndroidAction", # a short description of this parameter
                                       verify_block: proc do |value|
                                          UI.user_error!("No API token for BuildGolangAndroidAction given, pass using `api_token: 'token'`") unless (value and not value.empty?)
                                          # UI.user_error!("Couldn't find file at path '#{value}'") unless File.exist?(value)
                                       end),
          FastlaneCore::ConfigItem.new(key: :development,
                                       env_name: "FL_BUILD_GOLANG_ANDROID_DEVELOPMENT",
                                       description: "Create a development certificate instead of a distribution one",
                                       is_string: false, # true: verifies the input is a string, false: every kind of value
                                       default_value: false) # the default value if the user didn't provide one
        ]
      end

      def self.output
        # Define the shared values you are going to provide
        # Example
        [
          ['BUILD_GOLANG_ANDROID_CUSTOM_VALUE', 'A description of what this value contains']
        ]
      end

      def self.return_value
        # If your method provides a return value, you can describe here what it does
      end

      def self.authors
        # So no one will ever forget your contribution to fastlane :) You are awesome btw!
        ["Your GitHub/Twitter Name"]
      end

      def self.is_supported?(platform)
        # you can do things like
        #
        #  true
        #
        #  platform == :ios
        #
        #  [:ios, :mac].include?(platform)
        #

        platform == :android
      end
    end
  end
end
