"""empty message

Revision ID: 0c0edb6fd89b
Revises: d5ba47fd26e9
Create Date: 2022-04-04 13:39:19.409917

"""
import sqlalchemy as sa
from alembic import op
from sqlalchemy.dialects import postgresql

# revision identifiers, used by Alembic.
revision = '0c0edb6fd89b'
down_revision = 'd5ba47fd26e9'
branch_labels = None
depends_on = None


def upgrade():
    """
    Run upgrade
    """
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_index('ix_domain_proxy_logs_cbsd_serial_number', table_name='domain_proxy_logs')
    op.drop_index('ix_domain_proxy_logs_created_date', table_name='domain_proxy_logs')
    op.drop_index('ix_domain_proxy_logs_fcc_id', table_name='domain_proxy_logs')
    op.drop_index('ix_domain_proxy_logs_response_code', table_name='domain_proxy_logs')
    op.drop_table('domain_proxy_logs')
    # ### end Alembic commands ###


def downgrade():
    """
    Run downgrade
    """
    # ### commands auto generated by Alembic - please adjust! ###
    op.create_table(
        'domain_proxy_logs',
        sa.Column('id', sa.INTEGER(), autoincrement=True, nullable=False),
        sa.Column('log_from', sa.VARCHAR(), autoincrement=False, nullable=True),
        sa.Column('log_to', sa.VARCHAR(), autoincrement=False, nullable=True),
        sa.Column('log_name', sa.VARCHAR(), autoincrement=False, nullable=True),
        sa.Column('log_message', sa.VARCHAR(), autoincrement=False, nullable=True),
        sa.Column('cbsd_serial_number', sa.VARCHAR(), autoincrement=False, nullable=True),
        sa.Column('network_id', sa.VARCHAR(), autoincrement=False, nullable=True),
        sa.Column('fcc_id', sa.VARCHAR(), autoincrement=False, nullable=True),
        sa.Column('response_code', sa.INTEGER(), autoincrement=False, nullable=True),
        sa.Column('created_date', postgresql.TIMESTAMP(timezone=True), server_default=sa.text('statement_timestamp()'), autoincrement=False, nullable=False),
        sa.PrimaryKeyConstraint('id', name='domain_proxy_logs_pkey'),
    )
    op.create_index('ix_domain_proxy_logs_response_code', 'domain_proxy_logs', ['response_code'], unique=False)
    op.create_index('ix_domain_proxy_logs_fcc_id', 'domain_proxy_logs', ['fcc_id'], unique=False)
    op.create_index('ix_domain_proxy_logs_created_date', 'domain_proxy_logs', ['created_date'], unique=False)
    op.create_index('ix_domain_proxy_logs_cbsd_serial_number', 'domain_proxy_logs', ['cbsd_serial_number'], unique=False)
    # ### end Alembic commands ###
