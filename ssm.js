const AWS = require('aws-sdk');

const ssm = new AWS.SecretsManager();

module.exports.AURORA_PASSWORD = () => {
    // todo generate password if doesn't exist

    return ssm.getSecretValue({SecretId: 'EVS_AURORA_PASSWORD'})
        .promise()
        .then(secret => secret.SecretString);
};