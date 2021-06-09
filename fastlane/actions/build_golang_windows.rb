module Fastlane
  module Actions
    module SharedValues
      BUILD_GOLANG_WINDOWS_CUSTOM_VALUE = :BUILD_GOLANG_WINDOWS_CUSTOM_VALUE
    end

    class BuildGolangWindowsAction < Action
      def self.run(params)
        # fastlane will take care of reading in the parameter and fetching the environment variable:
        # UI.message "Parameter API Token: #{params[:api_token]}"

        if !File.directory?("build/")
          sh "mkdir build/"
        end
        sh "env GOOS=windows GOARCH=amd64 go build -o PortScanner-windows-amd64.exe -ldflags='-H windowsgui' . "
        sh "mv PortScanner-windows-amd64.exe build/"
      
        sh "env GOOS=windows GOARCH=arm go build -o PortScanner-windows-arm.exe -ldflags='-H windowsgui' ."
        sh "mv PortScanner-windows-arm.exe build/"

        sh "cd build/"
        sh "env GOOS=linux GOARCH=amd64 go build -o PortScanner-linux-amd64 ../main.go"

        # Actions.lane_context[SharedValues::BUILD_GOLANG_WINDOWS_CUSTOM_VALUE] = "my_val"
      end

      #####################################################
      # @!group Documentation
      #####################################################

      def self.description
        "A short description with <= 80 characters of what this action does"
      end

      def self.details
        # Optional:
        # this is your chance to provide a more detailed description of this action
        "You can use this action to do cool things..."
      end

      def self.available_options
        # Define all options your action supports.

        # Below a few examples
        []
      end

      def self.output
        # Define the shared values you are going to provide
        # Example
        [
          ['BUILD_GOLANG_WINDOWS_CUSTOM_VALUE', 'A description of what this value contains']
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

        platform == :ios
      end
    end
  end
end
